import Players from "./Players/Players";
import Menu from "./Menu/Menu";
import Main from "./Main/Main";
import Header from "./Header/Header";
import React, {useState, useEffect} from 'react'
import './App.css'

const positions = ["All", "Goalkeeper", "Defender", "Midfielder", "Forward"]
const filter = ["Price", "Total points", "Goals", "Assists"]
const values = [0,1,2,3,4,5,6,7,8,9,10]
const categories = ["Fantasy", "Supersport HNL", "Forum", "Draft", "Profile"]

function App() {
  const [position, setPosition] = useState("All")
  const [playerIds, setPlayerIds] = useState([])
  const [team, setTeam] = useState({
    'goalkeeper'  :[],
    'defenders'   :[],
    'midfielders' :[],
    'forwards'    :[]   
  })

  function addPlayer(position, player, id) {
    switch (position) {
      case 'G':
        position = 'goalkeeper'
        break
      case 'D':
        position = 'defenders'
        break
      case 'M':
        position = 'midfielders'
        break
      case 'F':
        position = 'forwards'
        break
    }

    let array = team[position].slice()
    array.push(player)
    setTeam({
      ...team,
      [position]: array
    })
    setPlayerIds([...playerIds, id])
  }

  function resetTeam() {
    setTeam({
      'goalkeeper'  :[],
      'defenders'   :[],
      'midfielders' :[],
      'forwards'    :[]   
    })
    setPlayerIds([])
  }
  return (
    <div className="player-section">
      <Menu categories={categories}/>
      <div className="header-and-field">
        <Header matchday="10. kolo" action={resetTeam} ids={playerIds}/>
        <Main changePosition={setPosition} buttonText={team}/>
      </div>
      <Players positions={positions} filter={filter} values={values} position={position} setPosition={setPosition} selectPlayer={addPlayer}/>
    </div>
  );
}

export default App;
