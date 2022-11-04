from bs4 import BeautifulSoup
from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.chrome.service import Service
import csv 
import requests

header = ['id', 'name', 'team', 'position', 'value']
data_file = open('players_database.csv', 'w', encoding="utf-8")
csv_writer = csv.writer(data_file, lineterminator='\n')
csv_writer.writerow(header)

selenium_driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()))
league_url = 'https://www.sofascore.com/tournament/football/croatia/hnl/170'
selenium_driver.get(league_url)

league_soup = BeautifulSoup(selenium_driver.page_source, 'lxml')
teams = league_soup.select("a[href*=team]")

teams_url = set()

for team in teams:
    teams_url.add('https://www.sofascore.com' + team['href'])

for team_url in teams_url:
    selenium_driver.get(team_url)
    total_height = int(selenium_driver.execute_script("return document.body.scrollHeight"))

    for i in range(1, total_height, 5):
        selenium_driver.execute_script("window.scrollTo(0, {});".format(i))
    team_soup = BeautifulSoup(selenium_driver.page_source, 'lxml')
    players = team_soup.select("a[href*=player]")

    player_ids = set()

    for player in players:
        href = player['href']
        id = href.split('/')[-1]
        player_ids.add(id)

    for id in player_ids:            
        req = requests.get('https://api.sofascore.com/api/v1/player/{}/'.format(id))
        req_json = req.json()
        
        player_id = id
        player_name = req_json['player']['name']
        player_team = req_json['player']['team']['name']
        if 'position' in req_json['player']:
            player_position = req_json['player']['position']
        else:
            continue
        if 'proposedMarketValue' in req_json['player']:
            player_value = req_json['player']['proposedMarketValue']
        else:
            player_value = 0

        player = [player_id, player_name, player_team, player_position, player_value]
        csv_writer.writerow(player)

data_file.close()