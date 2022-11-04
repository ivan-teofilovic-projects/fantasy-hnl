import "./Main.css"
import Position from "./Position";
import Field from "./Field";

const formation = [1,3,4,3]

function Main(props) {
    return (
        <div className="main">
            <Field />
            <div className="grouped-sections">
                <Position type="Goalkeeper" nplayers={formation[0]} changePosition={props.changePosition} buttonText={props.buttonText.goalkeeper}/>
                <Position type="Defender" nplayers={formation[1]} changePosition={props.changePosition} buttonText={props.buttonText.defenders}/>
                <Position type="Midfielder" nplayers={formation[2]} changePosition={props.changePosition} buttonText={props.buttonText.midfielders}/>
                <Position type="Forward" nplayers={formation[3]} changePosition={props.changePosition} buttonText={props.buttonText.forwards}/>
            </div>


        </div>
    )
}

export default Main;