import "./Header.css"
import Matchday from "./Matchday";
import TeamInfo from "./TeamInfo";

function Header(props) {
    return(
        <div className="Header">
            <Matchday matchday={props.matchday}/>
            <TeamInfo action={props.action} ids={props.ids}/>
        </div>
    );
}

export default Header;