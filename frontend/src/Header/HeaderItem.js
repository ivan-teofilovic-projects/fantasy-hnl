import "./HeaderItem.css"

function HeaderItem(props) {
    return (
        <div className="info-item">
            <span>{props.current}/{props.total}</span>
            <button onClick= {() => props.action()}>{props.buttonText}</button>
        </div>
    );
}

export default HeaderItem;