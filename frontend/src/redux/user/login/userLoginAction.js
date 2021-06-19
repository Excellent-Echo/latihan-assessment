import bookListAPI from "../../../APIs/booklistapi";

const resetForm = () => {
  return {
    type: "USER_LOGIN_RESET_FORM",
  };
};

const setEmail = email => {
  return {
    type: "USER_LOGIN_SET_EMAIL",
    payload: {
      email: email,
    },
  };
};

const setPassword = password => {
  return {
    type: "USER_LOGIN_SET_PASSWORD",
    payload: {
      password: password,
    },
  };
};

const setErrorMessage = errorMessage => {
  return {
    type: "USER_LOGIN_SET_ERROR_MESSAGE",
    payload: {
      errorMessage: errorMessage,
    },
  };
};

const startLoading = () => {
  return {
    type: "USER_LOGIN_START_LOADING",
  };
};

const stopLoading = () => {
  return {
    type: "USER_LOGIN_STOP_LOADING",
  };
};

const login = (email, password) => async dispatch => {
  try {
    dispatch(startLoading())
    dispatch(setErrorMessage(""))
    const submitData = {
      email: email,
      password: password,
    };

    const user = await bookListAPI({
      method: "POST",
      url: "/users/login",
      data: submitData,
    });

    const accessToken = user.data.token;
    localStorage.setItem("accessToken", accessToken);

    dispatch(stopLoading());
  } catch (error) {
    console.log(error);
    dispatch(setErrorMessage(error.response.data.error || ["internal server error"]));
    dispatch(stopLoading());
  }
}

const userLoginAction = {
  resetForm,
  setEmail,
  setPassword,
  login,
};

export default userLoginAction;
