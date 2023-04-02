
import Link from 'next/link'
function GridList({ games }) {
  return (
    <div className="grid grid-cols-4 gap-4 sm:grid-cols-4">
      {games.map((game: any) => (
        <Link href={`/games/${game.id}/gameflow`}>
          <div
            key={game.id}
            className="relative flex items-center space-x-3 rounded-lg border border-gray-300 bg-white px-6 py-5 shadow-sm focus-within:ring-2 focus-within:ring-indigo-500 focus-within:ring-offset-2 hover:border-gray-400"
          >
            <div className="flex-shrink-0">
              <img className="h-10 w-10 rounded-full" src={game.home_team_icon} alt="" />
            </div>
            <div className="min-w-0 flex-1">
              <a href="#" className="focus:outline-none">
                <span className="absolute inset-0" aria-hidden="true" />
                <p className="text-sm font-medium text-gray-900">{game.home_team}</p>
                <p className="truncate text-sm text-gray-500">{game.home_team_score}</p>
              </a>
            </div>
            <div className="flex-shrink-0">
              <img className="h-10 w-10 rounded-full" src={game.away_team_icon} alt="" />
            </div>
            <div className="min-w-0 flex-1">
              <a href="#" className="focus:outline-none">
                <span className="absolute inset-0" aria-hidden="true" />
                <p className="text-sm font-medium text-gray-900">{game.away_team}</p>
                <p className="truncate text-sm text-gray-500">{game.away_team_score}</p>
              </a>
            </div>
          </div>
        </Link>
      ))}
    </div>
  )
}

export default GridList;