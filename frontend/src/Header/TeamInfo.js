import "./TeamInfo.css"
import HeaderItem from "./HeaderItem";
import React, {useState, useEffect} from 'react'

function TeamInfo(props) {
    function handleClick() {
        fetch('http://127.0.0.1:8083', {
            method: 'POST',
            body: JSON.stringify(props.ids)
        }).then(response => {
            return(response.json())
            
        }).then(data => {
            alert("You have scored: " + data)
        })
        console.log(JSON.stringify(props.ids))
    }
    
    return (
        <div className="team-info">
            <HeaderItem current="0" total="15" buttonText="Auto Pick"/>
            <button onClick={() => handleClick()}>Submit</button>
            <HeaderItem current="0" total="100" buttonText="Reset" action={props.action}/>
        </div>
    )
}

export default TeamInfo;