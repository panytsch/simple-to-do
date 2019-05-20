import React from "react"

import {Todo} from "../redux/structs"

export const TodoComponent :React.FC = (props :any) => {
    const todo :Todo = props.Todo;
    return (
        <div className="todo" key={props.key}>
            <span>{todo.ID}</span>
            <input type="checkbox" checked={todo.IsDone} onChange={props.handlerDone}/>
            <span>{todo.Title}</span>
            <span>{todo.Description}</span>
            <button onClick={props.handlerDelete}>Delete</button>
        </div>
    )
};
