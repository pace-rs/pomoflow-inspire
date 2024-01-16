'use client';

import ThemeSwitcher from '@/components/theme/theme-switcher';
import { useRouter } from 'next/navigation';

export default function Home() {
  const router = useRouter();
  return (
    <main className='flex min-h-screen flex-col items-center justify-between p-24'>
      Pomotimer
      <ThemeSwitcher />
    </main>
  );
}
