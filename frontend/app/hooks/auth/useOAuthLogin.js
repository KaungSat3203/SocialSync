export const useOAuthLogin = () => {
  const baseUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

  const loginWithFacebook = () => {
    window.location.href = `${baseUrl}/auth/facebook/login`;
  };

  const loginWithGoogle = () => {
    window.location.href = `${baseUrl}/auth/google/login`;
  };

  return { loginWithFacebook, loginWithGoogle };
};
