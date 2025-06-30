// app/layout.js
import localFont from 'next/font/local';
import './globals.css';
import { Toaster } from 'react-hot-toast';

const geistSans = localFont({
  src: [
    {
      path: './fonts/Geist/Geist-Regular.woff2',
      weight: '400',
      style: 'normal',
    },
    {
      path: './fonts/Geist/Geist-Bold.woff2',
      weight: '700',
      style: 'normal',
    },
  ],
  variable: '--font-geist-sans',
});

const geistMono = localFont({
  src: [
    {
      path: './fonts/Geist/GeistMono-Regular.woff2',
      weight: '400',
      style: 'normal',
    },
  ],
  variable: '--font-geist-mono',
});

export const metadata = {
  title: 'SocialSync',
  description: 'Collaborative social media management tool',
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
        <Toaster
          position="bottom-right"
          toastOptions={{
            style: {
              fontSize: '0.9rem',
              background: '#fff',
              color: '#333',
              border: '1px solid #e5e7eb',
            },
          }}
        />
        {children}
      </body>
    </html>
  );
}
