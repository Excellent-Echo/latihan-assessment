const resetForm = () =>{
    return{
        type:"USER_REISTER_RESET_FORM"
    };
};
const setName = Name =>{
    return{
        type:"USER_SET_NAME",
        payload:{
            Name:Name,
        },
    };
};
const setEmail = Email =>{
    return{
        type:"USER_SET_EMAIL",
        payload:{
            Email:Email,
        },
    };
};

const setPassword = Password =>{
    return{
        type:"USER_SET_PASSWORD",
        payload:{
            Password:Password,
        },
    };
};

const setAddress = Address =>{
    return{
        type:"USER_SET_ADDRESS",
        payload:{
            Address:Address,
        },
    };
};

const setDateBirth = DateBirth =>{
    return{
        type:"USER_SET_DATEBIRTH",
        payload:{
            DateBirth:DateBirth,
        },
    };
};