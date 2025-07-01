'use client';

import { useCallback, useEffect, useState } from 'react';
import axios from 'axios';
import SocialAccountCard from '../../components/SocialAccountCard';
import DisconnectModal from '../../components/DisconnectModal';

import { SiMastodon } from 'react-icons/si';
import {
  FaFacebook,
  FaInstagram,
  FaYoutube,
  FaTiktok,
  FaTwitter,
  FaTelegram,
} from 'react-icons/fa';
import { FaSquareThreads } from 'react-icons/fa6';

const backendToDisplayName = {
  facebook: 'Facebook',
  instagram: 'Instagram',
  youtube: 'YouTube',
  twitter: 'Twitter (X)',
  mastodon: 'Mastodon',
  threads: 'Threads',
  telegram: 'Telegram',
  tiktok: 'TikTok',
};

const platformsList = [
  'facebook',
  'instagram',
  'youtube',
  'twitter',
  'mastodon',
  'threads',
  'telegram',
  'tiktok',
];

const getIcon = (platform) => {
  switch (platform) {
    case 'Facebook':
      return FaFacebook;
    case 'Instagram':
      return FaInstagram;
    case 'YouTube':
      return FaYoutube;
    case 'TikTok':
      return FaTiktok;
    case 'Twitter (X)':
      return FaTwitter;
    case 'Mastodon':
      return SiMastodon;
    case 'Threads':
      return FaSquareThreads;
    case 'Telegram':
      return FaTelegram;
    default:
      return null;
  }
};

export default function ManageAccountPage() {
  const [platforms, setPlatforms] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [statusMessage, setStatusMessage] = useState('');
  const [statusType, setStatusType] = useState('success');
  const baseUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

  const [showConfirmModal, setShowConfirmModal] = useState(false);
  const [platformToDisconnect, setPlatformToDisconnect] = useState(null);

  const token = typeof window !== 'undefined' ? localStorage.getItem('accessToken') : null;

  const fetchAccounts = useCallback(async () => {
    if (!token) {
      setError('Access token not found.');
      setLoading(false);
      return;
    }

    try {
      const res = await axios.get(`${baseUrl}/api/social-accounts`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const accounts = Array.isArray(res.data) ? res.data : [];

      const platformData = platformsList.map((backendKey) => {
        const displayName = backendToDisplayName[backendKey];
        const account = accounts.find(
          (acc) => acc?.platform?.toLowerCase() === backendKey
        );

        return {
          name: displayName,
          icon: getIcon(displayName),
          connected: !!account,
          userProfilePic:
            account?.profilePictureUrl && account.profilePictureUrl !== 'null'
              ? account.profilePictureUrl
              : null,
          accountName: account?.profileName || '',
        };
      });

      setPlatforms(platformData);
      setError(null);
    } catch (err) {
      console.error(err);
      setError('Failed to fetch social accounts.');
    } finally {
      setLoading(false);
    }
  }, [token]);

  useEffect(() => {
    fetchAccounts();
  }, [fetchAccounts]);

  useEffect(() => {
    if (statusMessage) {
      const timer = setTimeout(() => setStatusMessage(''), 5000);
      return () => clearTimeout(timer);
    }
  }, [statusMessage]);

  const handleConnect = async (platformName, isConnected) => {
    if (!token) {
      setStatusMessage('You must be logged in.');
      setStatusType('error');
      return;
    }

    const isFacebookConnected = platforms.find((p) => p.name === 'Facebook')?.connected;

    if (!isConnected) {
      if (platformName === 'Instagram' && !isFacebookConnected) {
        setStatusMessage('Please connect your Facebook Page first before connecting Instagram.');
        setStatusType('error');
        return;
      }

      try {
        const urlMap = {
          Facebook: `${baseUrl}/auth/facebook/login?token=${token}`,
          YouTube: `${baseUrl}/auth/youtube/login?token=${token}`,
          TikTok: `${baseUrl}/auth/tiktok/login?token=${token}`,
          'Twitter (X)': `${baseUrl}/auth/twitter/login?token=${token}`,
          Threads: `${baseUrl}/auth/threads/login?token=${token}`,
          Telegram: `${baseUrl}/auth/telegram/login?token=${token}`,
        };

        if (platformName === 'Instagram') {
          await axios.post(
            `${baseUrl}/connect/instagram`,
            {},
            { headers: { Authorization: `Bearer ${token}` } }
          );
          setStatusMessage('Instagram account connected successfully!');
          setStatusType('success');
          fetchAccounts();
        } else if (platformName === 'Mastodon') {
          const instance = 'mastodon.social';
          window.location.href = `${baseUrl}/auth/mastodon/login?instance=${encodeURIComponent(
            instance
          )}&token=${token}`;
        } else if (urlMap[platformName]) {
          window.location.href = urlMap[platformName];
        } else {
          setStatusMessage(`Connect to ${platformName} is not yet implemented.`);
          setStatusType('error');
        }
      } catch (err) {
        const msg = err?.response?.data?.error || `Failed to connect ${platformName}.`;
        setStatusMessage(msg);
        setStatusType('error');
      }
    } else {
      setPlatformToDisconnect(platformName);
      setShowConfirmModal(true);
    }
  };

  const handleConfirmDisconnect = async () => {
    setShowConfirmModal(false);
    if (!platformToDisconnect) return;

    try {
      await axios.delete(
        `${baseUrl}/api/social-accounts/${platformToDisconnect.toLowerCase()}`,
        { headers: { Authorization: `Bearer ${token}` } }
      );

      setPlatforms((prev) =>
        prev.map((p) =>
          p.name === platformToDisconnect
            ? { ...p, connected: false, userProfilePic: null, accountName: '' }
            : p
        )
      );

      setStatusMessage(`${platformToDisconnect} disconnected successfully.`);
      setStatusType('success');
    } catch (err) {
      const msg = err?.response?.data?.error || `Failed to disconnect ${platformToDisconnect}.`;
      setStatusMessage(msg);
      setStatusType('error');
    } finally {
      setPlatformToDisconnect(null);
    }
  };

  const handleCloseConfirmModal = () => {
    setShowConfirmModal(false);
    setPlatformToDisconnect(null);
  };

  return (
    <div className="p-6">
      <div className="flex items-center mb-6 border-b pb-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-800">Hello!</h1>
          <p className="text-gray-600">Manage Your Social Media Accounts</p>
        </div>
      </div>

      {statusMessage && (
        <div
          className={`mb-4 p-3 rounded-lg text-sm ${
            statusType === 'success'
              ? 'bg-green-100 text-green-800'
              : 'bg-red-100 text-red-800'
          }`}
          role="alert"
          aria-live="polite"
        >
          {statusMessage}
        </div>
      )}

      {loading ? (
        <p className="text-gray-500">Loading...</p>
      ) : error ? (
        <p className="text-red-500">{error}</p>
      ) : (
        <>
          <p className="mb-6 text-gray-600">
            Connect or disconnect your accounts to start managing content across platforms.
          </p>

          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
            {platforms.map((platform) => (
              <SocialAccountCard
                key={platform.name}
                platform={platform.name}
                IconComponent={platform.icon}
                connected={platform.connected}
                userProfilePic={platform.userProfilePic}
                accountName={platform.accountName}
                onConnect={() => handleConnect(platform.name, platform.connected)}
              />
            ))}
          </div>
        </>
      )}

      <DisconnectModal
        show={showConfirmModal}
        onClose={handleCloseConfirmModal}
        onConfirm={handleConfirmDisconnect}
        platformName={platformToDisconnect}
      />
    </div>
  );
}
