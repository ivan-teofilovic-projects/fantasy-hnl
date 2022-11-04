import Filter from "./Filter";
import Player from "./Player";
import Search from "./Search";
import React, {useState, useEffect} from 'react'

function Players(props) {
    const [players, setPlayers] = useState([])
    const [searchTerm, setSearchTerm] = useState("")
    const [sort, setSort] = useState("Price")
    const [value, setValues] = useState(10)
    useEffect(() => {
        fetch('http://127.0.0.1:8080')
        .then(response => {
            return response.json()
            
        }
        ).then(data => {
            setPlayers(data)
        }
        )
    })
    return (
        <div>
            <Search handleType={setSearchTerm}/>
            <Filter type="Position" items={props.positions} handleSelect={props.setPosition}/>
            <Filter type="Sort" items={props.filter} handleSelect={setSort}/>
            <Filter type="Value" items={props.values} handleSelect={setValues}/>
            <Player players={players} searchTerm={searchTerm} position={props.position} sort={sort} value={value} selectPlayer={props.selectPlayer}/>
        </div>

    );
}

export default Players;