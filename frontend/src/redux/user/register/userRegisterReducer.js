const initialState = {
    Name:"",
    Email:"",
    Password:"",
    Address:"",
    DateBirth:""
};

const userRegisterReducer = (state  = initialState, action)=>{
    switch (action.type){
        case "USER_REGISTER_RESET_FORM":
            return {
                ...initialState,
            };
        case "USER_SET_NAME":
            return{
                ...state,
                Name:action.payload.Name,
            };
        case "USER_SET_EMAIL":
            return{
                ...state,
                Email:action.payload.Email,
            };
        case "USER_SET_PASSWORD":
            return{
                ...state,
                Password:action.payload.Password,
            };
        case "USER_SET_ADDRESS":
            return{
                ...state,
                Address:action.payload.Address,
            };
        case "USER_SET_DATEBIRTH":
            return{
                ...state,
                Name:action.payload.DateBirth,
            };

        default:
            return state;
    }
}

export default userRegisterReducer;