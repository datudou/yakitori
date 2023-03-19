'use client';
import { Inter } from 'next/font/google'
import Navbar from '@/components/Navbar'
import Table from '@/components/Table'
import { useState, useEffect } from 'react';



const inter = Inter({ subsets: ['latin'] })

export default function Home() {

  return (
    <main>
      <div className="min-h-full">
        <Navbar></Navbar>
        <div className="py-10">
          <header>
          </header>
          <main>
            <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">
              <Table></Table>
            </div>
          </main>
        </div>
      </div>
    </main>
  )
}
