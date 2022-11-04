import "./Position.css"
import kit from './images/kit.png'
import Player from "./Player";

function Position(props) {
    const nplayers = props.nplayers
    const positions = Array.from({length: nplayers}, (_, i) => {
        return (
            <div>
                <img src={kit}></img>
                <Player type={props.type} changePosition={props.changePosition} player={props.buttonText[i]}/>
            </div>
        )
    });
    return (
        <div className="positionss">
            {positions}
        </div>
    );
}

export default Position;