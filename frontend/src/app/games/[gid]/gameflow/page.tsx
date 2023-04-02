import { Inter } from 'next/font/google'
import Navbar from '@/components/Navbar'
import _ from 'lodash';


import dynamic from 'next/dynamic';

const Chart = dynamic(
  () => import('@/components/Chart'),
  { ssr: false }
);

async function fetchGameLogs(gid: string) {
  const res = await fetch(`http://localhost:8080/api/v1/game/${gid}/gamelog`);
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
                <div>
                  <Chart isHome={true} gameLogs={homeTeamGameLogs} quater={1} isShowYAxis={true} />
                </div>
                <div>
                  <Chart isHome={true} gameLogs={homeTeamGameLogs} quater={2} isShowYAxis={false} />
                </div>
                <div>
                  <Chart isHome={true} gameLogs={homeTeamGameLogs} quater={3} isShowYAxis={false} />
                </div>
                <div>
                  <Chart isHome={true} gameLogs={homeTeamGameLogs} quater={4} isShowYAxis={false} />
                </div>
              </div>
            </div>
            <div className="mx-auto max-w-7xl sm:px-6 lg:px-8">
              <div className="grid lg:grid-cols-4 gap-0">
                <div>
                  <Chart isHome={false} gameLogs={awayTeamGameLogs} quater={1} isShowYAxis={true} />
                </div>
                <div >
                  <Chart isHome={false} gameLogs={awayTeamGameLogs} quater={2} isShowYAxis={false} />
                </div>
                <div >
                  <Chart isHome={false} gameLogs={awayTeamGameLogs} quater={3} isShowYAxis={false} />
                </div>
                <div >
                  <Chart isHome={false} gameLogs={awayTeamGameLogs} quater={4} isShowYAxis={false} />
                </div>
              </div>
            </div>


          </main>
        </div>
      </div>
    </main>
  )
}
