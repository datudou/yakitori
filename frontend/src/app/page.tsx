'use client';
import { Inter } from 'next/font/google'
import Navbar from '../components/Navbar'
import GridList from '../components/GridList'
import { useState, useEffect } from 'react';



const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const [games, setGames] = useState([]);
  useEffect(() => {
    async function fetchGames() {
      try {
        const resp = await fetch("/api/v1/games/2023-03-16")
        const data = await resp.json()
        setGames(await data)
      } catch (error) {
        console.error(error)
      }
    }
    fetchGames()
  }, [])

  return (
    <main>
      <div className="min-h-full">
        <Navbar></Navbar>
        <div className="py-10">
          <header>
            <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
              <h1 className="text-3xl font-bold leading-tight tracking-tight text-gray-900">Dashboard</h1>
            </div>
          </header>
          <main>
            <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">
              <GridList games={games}></GridList>
            </div>
          </main>
        </div>
      </div>
    </main>
  )
}
