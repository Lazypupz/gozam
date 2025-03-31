
function UserGreeting(props){
    if(props.isLoggedIn){
        return <h2>Welcome {props.UserName}</h2>
    }
    else {
        return <h2>Please Log in to continue</h2>
    }
}

export default UserGreeting