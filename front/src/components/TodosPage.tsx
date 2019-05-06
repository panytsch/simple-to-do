import React from "react";
import {WsHost} from "../redux/methods";
import {ReduxState, Todo} from "../redux/structs";
import {connect} from "react-redux";
import {withRouter} from "react-router";
import {TodoComponent} from "./Todo";
import {EventListener, WsEvent, WsRequest} from "../helpers/Ws";

class TodosPage extends React.Component<any> {
    readonly WS :WebSocket;
    private newTodoTitle :string = '';
    private newTodoDescription :string = '';

    constructor(props :any) {
        super(props);
        if (props.Token === '') {
            props.history.push('/') //redirect
        }
        this.WS = new WebSocket(`${WsHost}/todo?token=${props.Token}`);
    }
    WsSend(request :WsRequest) :void {
        this.WS.send(JSON.stringify(request));
    }
    componentDidMount(): void {
        this.WS.addEventListener("message", EventListener);
        this.WS.onopen = () => {
            this.WsSend({
                Type: WsEvent.Connect,
                Token: this.props.Token
            });
        };
        this.WS.onclose = () => {
            this.props.history.push('/');
        };
    }
    componentWillUnmount(): void {
        if (this.WS !== null) {
            this.WS.close();
        }
    }

    AddTodo(e :Event) {
        e.preventDefault();
        const NewTodo :Todo = {
            Title: this.newTodoTitle,
            Description: this.newTodoDescription,
        };
        const WsRequest :WsRequest = {
            Token: this.props.Token,
            Type: WsEvent.Add,
            Todo: NewTodo,
        };
        this.WsSend(WsRequest);
    }

    render(): React.ReactElement<any, string | React.JSXElementConstructor<any>> | string | number | {} | React.ReactNodeArray | React.ReactPortal | boolean | null | undefined {
        const {Todos} = this.props;
        return <div>
            <form>
                <label htmlFor="new_title">Title</label>
                <input type="text" placeholder="title" id="new_title" onChange={e => this.newTodoTitle = e.target.value}/>

                <label htmlFor="new_desc">Title</label>
                <input type="text" placeholder="title" id="new_desc" onChange={e => this.newTodoDescription = e.target.value}/>

                {
                    // @ts-ignore
                    <button onClick={(e: Event) => this.AddTodo(e)}>Add</button>
                }
            </form>
            <h3>My todos</h3>
            {
                Todos.map((todo :Todo) => {
                    // @ts-ignore
                    return <TodoComponent Todo={todo}/>
                })
            }
        </div>;
    }
}


const mapDispatchToProps = () => ({

});

const mapStateToProps = (state :any) :ReduxState => state.data;

export default connect(mapStateToProps, mapDispatchToProps)(
    withRouter(TodosPage)
);