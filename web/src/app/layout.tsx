import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

import { Inter as FontSans } from "next/font/google";

import { cn } from "@/lib/utils";
import { ReactNode } from "react";
import { ThemeProvider } from "./providers";

const fontSans = FontSans({
  subsets: ["latin"],
  variable: "--font-sans",
});

const inter = Inter({ subsets: ["latin"] });

interface RootLayoutProperties {
  children: ReactNode;
}

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({ children }: RootLayoutProperties) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={`cn("min-h-screen bg-background font-sans antialiased", fontSans.variable)`}>
        <ThemeProvider attribute="class" defaultTheme="light" enableSystem disableTransitionOnChange>
          <div className="mt-4 flex min-h-[calc(100vh-185px)] flex-col">
            <div className="container mx-auto flex-grow py-5">{children}</div>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}
