import "./Field.css"
import field from "./images/Football_field.png"
function Field(props) {
    return (
        <div className="field">
            <img src={field}></img>
        </div>
    );
}

export default Field;