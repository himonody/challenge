import type { Metadata } from "next";
import { ReactNode } from "react";
import { Inter } from "next/font/google";
import { NextIntlClientProvider } from "next-intl";
import "../globals.css";
import { Providers } from "../providers";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Challenge App",
  description: "Challenge frontend",
};

type LayoutProps = {
  children: ReactNode;
  params: { locale: string };
};

async function getMessages(locale: string) {
  try {
    return (await import(`../../messages/${locale}.json`)).default;
  } catch (error) {
    return (await import("../../messages/en.json")).default;
  }
}

export default async function RootLayout({ children, params }: LayoutProps) {
  const locale = params.locale || "en";
  const messages = await getMessages(locale);
  return (
    <html lang={locale}>
      <body className={inter.className}>
        <NextIntlClientProvider locale={locale} messages={messages}>
          <Providers>{children}</Providers>
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
