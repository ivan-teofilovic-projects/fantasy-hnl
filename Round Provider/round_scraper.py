import requests

rounds = requests.get("https://api.sofascore.com/api/v1/unique-tournament/170/season/42138/rounds").json()
current_round = rounds['currentRound']['round']
round_url = "https://api.sofascore.com/api/v1/unique-tournament/170/season/42138/events/round/{}".format(current_round)
round = requests.get(round_url).json()
events = round['events']
previous_round = current_round - 1
for event in events:
    print(event['homeTeam']['name'] + " - " + event['awayTeam']['name'])