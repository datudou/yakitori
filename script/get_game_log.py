#!/usr/bin/python3
from basketball_reference_web_scraper import client
from basketball_reference_web_scraper.data import Team, OutputType
from datetime import datetime,timezone,timedelta
from time import sleep


def get_season_game(season_end_year):
    season_game = client.season_schedule(season_end_year=season_end_year)
    return season_game


team_name_map = {
    "Atlanta Hawks": "ATL",
    "Boston Celtics": "BOS",
    "Brooklyn Nets": "BKN",
    "Charlotte Hornets": "CHA",
    "Chicago Bulls": "CHI",
    "Cleveland Cavaliers": "CLE",
    "Dallas Mavericks": "DAL",
    "Denver Nuggets": "DEN",
    "Detroit Pistons": "DET",
    "Golden State Warriors": "GSW",
    "Houston Rockets": "HOU",
    "Indiana Pacers": "IND",
    "Los Angeles Clippers": "LAC",
    "Los Angeles Lakers": "LAL",
    "Memphis Grizzlies": "MEM",
    "Miami Heat": "MIA",
    "Milwaukee Bucks": "MIL",
    "Minnesota Timberwolves": "MIN",
    "New Orleans Pelicans": "NOP",
    "New York Knicks": "NYK",
    "Oklahoma City Thunder": "OKC",
    "Orlando Magic": "ORL",
    "Philadelphia 76ers": "PHI",
    "Phoenix Suns": "PHX",
    "Portland Trail Blazers": "POR",
    "Sacramento Kings": "SAC",
    "San Antonio Spurs": "SAS",
    "Toronto Raptors": "TOR",
    "Utah Jazz": "UTA",
    "Washington Wizards": "WAS",
}

def convert_teamname_to_simplename(team_name):
    new_team_map = {}
    for key,value in team_name_map.items():
        new_team_map[key.upper()] = value

    return new_team_map[team_name.value]


def get_game_log(date, home_team, away_team, home_team_score, away_team_score):
    dat = date.astimezone(timezone(timedelta(hours=-4),'America/New_York'))
    year = dat.year
    month = dat.month
    day = dat.day
    
    client.play_by_play(
        home_team= home_team,
        year=year, month=month, day=day,
        output_type=OutputType.JSON,
        output_file_path="./{date}|{home_team}-{home_team_score}vs{away_team}-{away_team_score}.json".\
            format(date=date ,\
                   home_team=convert_teamname_to_simplename(home_team),\
                   home_team_score=home_team_score,\
                   away_team=convert_teamname_to_simplename(away_team),\
                   away_team_score=away_team_score))




if __name__ == '__main__':
    games = get_season_game(2023)
    for game in games:
        date = game['start_time']
        print(game['home_team_score'],game['away_team_score'])
        
        sleep(10)

        if date < datetime.today().astimezone(timezone.utc):
            get_game_log(date, game['home_team'], game['away_team']\
                         ,game['home_team_score'], game['away_team_score'])
