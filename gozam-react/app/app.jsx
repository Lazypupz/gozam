import Student  from "./student.jsx";


function App(){
  return (
    <>
      <Student name="Spongebob" age={10} isStudent={true}></Student>
  
      <UserGreeting isLoggedIn={true} UserName="Lpz"></UserGreeting>
    
    </>
  );
}
export default App