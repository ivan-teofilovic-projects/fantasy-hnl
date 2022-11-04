import './Player.css'

function Player(props) {
    const players = props.players;
    const filtered_players = players.filter((player) => {
        if (props.searchTerm == "") {
            return player
        } else if (player.name.includes(props.searchTerm)){
            return player
        }
    }).filter((player) => {
        if (props.position == 'All'){
            return player
        } else if (props.position[0] == player.position) {
            return player
        }
    }).sort((a,b) => (parseInt(a.value) < parseInt(b.value)) ? 1 : ((parseInt(b.value) < parseInt(a.value)) ? -1 : 0))
    const listItems = filtered_players.map((player) => 
        <div className="player-info" onClick={() => {props.selectPlayer(player.position, player.name, player.id)}}>
            <div className="general-info">
                <div key={player.id} playerid={player.id}>{player.name}</div>
                <div>{player.position}</div>
            </div>
            <div className="value">
                <div>{player.value}</div>
            </div>
        </div>


        
    )
    return (
        <div className='listOfPlayers'>
            {listItems}
        </div>
        
    );
}

export default Player;