document.getElementById("signupForm").addEventListener("submit",async(e)=>{
  e.preventDefault();

console.log("🔥 NEW SIGNUP JS LOADED");

  const email=document.getElementById("email").value.trim();
  const password=document.getElementById("password").value;
  const confirmPassword=document.getElementById("confirmPassword").value;

  //basic validations

  if(password!==confirmPassword){
    alert("Password mismatch");
    return
  }

  if(password.length<6){
    alert("password must be atleast 6 characters");
    return;
  }
  try{
    const res=await fetch("http://localhost:8080/register",{
      method:"POST",
      headers:{
        "Content-Type":"application/json"
      },
      body:JSON.stringify({
        email:email,
        password:password
      })
    });
    let data;
    try{
      data = await res.json();

    }catch{
      data={};
    }
    if (res.ok){
      alert("Account created");
      window.location.href="index.html";

    }else{
      alert(data.error|| data.message||"signup failed");
    }
  }catch(err){
    console.error(err);
    alert("Server error");
  }
})