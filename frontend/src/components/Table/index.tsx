
import Bar from '../Bar'
interface Todo {
  playerName: string;
  onCourtTime: number;
}



export default function Table({ awayTeamGameLogs, homeTeamGameLogs }) {
  return (
    <div className="px-4 sm:px-6 lg:px-8">
      <div className="sm:flex sm:items-center">
        <div className="sm:flex-auto">
          <h1 className="text-base font-semibold leading-6 text-gray-900">Users</h1>
          <p className="mt-2 text-sm text-gray-700">
            A list of all the users in your account including their name, title, email and role.
          </p>
        </div>
      </div>
      <div className="mt-8 flow-root">
        <div className="overflow-x-auto ">
          <table>
            <thead className="">
              <tr className="divide-x divide-white bg-gray-200">
                <th scope="col" className="w-1/5 py-3.5 pl-4 pr-4 text-left text-sm font-semibold text-gray-900 sm:pl-0">
                  Player
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  1st
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  2nd
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  3rd
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  4th
                </th>
                <th scope="col" className="text-center w-1/5 px-4 py-3.5 text-left text-sm font-semibold text-gray-900">
                  Min
                </th>
              </tr>
            </thead>
            <tbody className="text-black  divide-y divide-x divid-white text-sm">
              {homeTeamGameLogs.map((gameLog: any) => (
                <tr>
                  <td className="p-0">
                    <div className="h-full w-full px-0 py-0">{gameLog.player_name}</div>
                  </td>
                  <td className="p-0">
                    {gameLog.mins[1] !== null ? <div className="h-full w-1/5 bg-blue-100 px-0 py-0">{calTimePercentOfPlayerOnTheCourt(gameLog, 1)}</div> : null}
                  </td>
                </tr>

              )
              )}
            </tbody>
            <tbody className="text-black divide-y divide-x text-sm divid-white">
              {awayTeamGameLogs.map((gameLog: any) => (
                <tr>
                  <td className="p-0">
                    <div className="h-full w-full  px-0 py-0">{gameLog.player_name}</div>
                  </td>
                  <td className="p-0 flex">

                    <Bar gameLog={gameLog} period={1}></Bar>
                    <div className="h-full w-1/2 bg-pink-400 px-0 py-0">xxx</div>
                    <div className="h-full w-1/2 bg-pink-900 px-0 py-0">222</div>
                  </td>
                </tr>
              )
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div >
  )

}










