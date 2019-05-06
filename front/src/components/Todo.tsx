import React from "react"

import {Todo} from "../redux/structs"

export const TodoComponent :React.FC = (props :any) => {
    const todo :Todo = props.Todo;
    return (
        <div className="todo">
            <span>{todo.ID}</span>
            <input type="checkbox" checked={todo.IsDone}/>
            <span>{todo.Title}</span>
            <span>{todo.Description}</span>
        </div>
    )
};
