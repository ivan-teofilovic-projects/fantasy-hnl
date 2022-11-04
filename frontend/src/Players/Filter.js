import "./Filter.css"

function Filter(props) {
    const items = props.items
    const listItems = items.map((item) => 
        <option value={item} key={item}>{item}</option>
    )
    return(
        <div className="filter">
            <label>{props.type}</label>
            <select name={props.type} id={props.type} onChange={(event) => {props.handleSelect(event.target.value)}}>
                {listItems}
            </select>
        </div>
    );
}

export default Filter;