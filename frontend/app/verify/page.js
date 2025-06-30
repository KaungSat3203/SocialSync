'use client';

import React, { Suspense } from 'react';
import VerifyEmailInner from './VerifyEmailInner';

const VerifyEmail = () => {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <VerifyEmailInner />
    </Suspense>
  );
};

export default VerifyEmail;
