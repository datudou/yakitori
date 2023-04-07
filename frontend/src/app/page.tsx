import Navbar from '../components/Navbar'
import GridList from '../components/GridList'



async function fetchGames() {
  const res = await fetch("/api/v1/games/2022-10-19");

  if (!res.ok) {
    throw new Error('Failed to fetch data');
  }

  return res.json();
}



export default async function Home() {
  const games = await fetchGames();

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
            <div className="mx-auto max-w-7xl sm:px-6 lg:px-8" >
              <GridList games={games}></GridList>
            </div>
          </main>
        </div>
      </div>
    </main>
  )
}
