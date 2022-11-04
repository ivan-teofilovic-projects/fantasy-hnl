import React from 'react'
import './Player.css'

function Player(props) {
    return (!props.player) ? (
        <button onClick={() => props.changePosition(props.type)}>Add Player</button>
    ) : <div className="text-overflow-center">{props.player}</div>
}

export default Player