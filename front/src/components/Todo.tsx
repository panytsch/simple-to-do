import React from "react"

import {Todo as TodoStruct} from "../redux/structs"

export const Todo :React.FC = (props :any) => {
    const todo :TodoStruct = props.Todo;
    return (
        <div className="todo">
            <span>{todo.ID}</span>
            <input type="checkbox" checked={todo.IsDone}/>
            <span>{todo.Title}</span>
            <span>{todo.Description}</span>
        </div>
    )
};