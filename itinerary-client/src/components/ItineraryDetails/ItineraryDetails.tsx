import React from "react";

export const ItineraryDetails = (props: {name: string, isCompleted: boolean}) => {
    console.log(props.name, props.isCompleted)
    return (
        <div>
            <h4>{props.name} = {props.isCompleted}</h4>
        </div>
    )
}