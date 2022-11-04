import Item from "./Item";
import "./Menu.css"

function Menu(props) {
    const categories = props.categories
    const categoryItems = categories.map((item) => 
        <Item category={item}/>
    )
    return(
        <div className="menu">
            {categoryItems}
        </div>
    );
}

export default Menu;