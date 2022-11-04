import "./Search.css"

function Search(props) {
    return(
        <input className="player-input" placeholder="Search for a player.." onChange={(event) => {props.handleType(event.target.value)}}></input>
    );
}

export default Search;