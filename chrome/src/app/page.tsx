'use client';

import ThemeSwitcher from '@/components/theme/theme-switcher';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';

export default function Home() {
  const router = useRouter();
  return (
    <main className='flex min-h-screen flex-col items-center justify-between p-24'>
      <ThemeSwitcher />
      <Button
        onClick={() => {
          router.push('/auth');
        }}
      >
        Log Out{' '}
      </Button>
    </main>
  );
}
