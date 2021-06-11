const initialState = {
  firstName: "",
  lastName: "",
  email: "",
  password: "",
};

const userRegisterReducer = (state = initialState, action) => {
  switch (action.type) {
    case "USER_REGISTER_FORM":
      return {
        ...initialState,
      };
    case "USER_REGISTER_SET_FIRSTNAME":
      return {
        ...state,
        firstName: action.payload.firstName,
      };
    default:
      return state;
  }
};

export default userRegisterReducer;
