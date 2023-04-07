import Navbar from '@/components/Navbar'


import dynamic from 'next/dynamic';

const Chart = dynamic(
  () => import('@/components/Chart'),
  { ssr: false }
);

async function fetchGameLogs(gid: string) {
  const res = await fetch(`/api/v1/game/${gid}/gamelog`);
  const data = await res.json();
  let homeTeamGameLogs = [];
  let awayTeamGameLogs = [];
  data.forEach((gameLog: any) => {
    if (gameLog.is_home_team) {
      homeTeamGameLogs.push(gameLog);
    } else {
      awayTeamGameLogs.push(gameLog);
    }
  });

  if (!res.ok) {
    throw new Error('Failed to fetch data');
  }
  return { homeTeamGameLogs, awayTeamGameLogs };
}

export default async function GameFlow({ params: { gid } }) {
  console.info(gid)

  const { homeTeamGameLogs, awayTeamGameLogs } = await fetchGameLogs(gid);
  console.info(homeTeamGameLogs)

  return (
    <main>
      <div className="min-h-full">
        <Navbar></Navbar>
        <div className="py-10">
          <header>
          </header>
          <main>
            <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">
              <div className="grid lg:grid-cols-4 gap-0">
                {[1, 2, 3, 4].map((quarter)=>
                <div>
                  <Chart isHome={true} gameLogs={homeTeamGameLogs} quarter={quarter} isShowYAxis={true} />
                </div>
                )}
              </div>
            </div>
            <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">
              <div className="grid lg:grid-cols-4 gap-0">
                {[1, 2, 3, 4].map((quarter)=>
                <div>
                  <Chart isHome={true} gameLogs={homeTeamGameLogs} quarter={quarter} isShowYAxis={true} />
                </div>
                )}
              </div>
            </div>
          </main>
        </div>
      </div>
    </main>
  )
}
