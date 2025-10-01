import Cookies from "js-cookie";

export const saveTokenFromURL = () => {
  const urlParams = new URLSearchParams(window.location.search);
  const dataStr = urlParams.get("data");

  if (!dataStr) return;

  try {
    const data = JSON.parse(decodeURIComponent(dataStr));
    const user_id = data.user.id;
    const token = data.access_token;

    if (token) {
      Cookies.set("access_token", token, { expires: 1, path: "/" });
      Cookies.set("user_id", user_id, { expires: 1, path: "/" });
      window.history.replaceState({}, document.title, window.location.pathname);
    }

    return { user_id, token };
  } catch (error) {
    console.error("Failed to parse data from URL:", error);
    return null;
  }
};

export const checkAuth = () => {
  const token = Cookies.get("access_token");
  const user_id = Cookies.get("user_id");

  if (!token || !user_id) {
    window.location.href = "/customers/login";
    return null;
  }

  return { token, user_id };
};

export const setupInactivityLogout = (timeoutMinutes = 3) => {
  let inactivityTimer;

  const resetTimer = () => {
    clearTimeout(inactivityTimer);
    inactivityTimer = setTimeout(() => {
      Cookies.remove("access_token");
      Cookies.remove("user_id");
      console.log("User inactive, logging out...");
      window.location.href = "/";
    }, timeoutMinutes * 60 * 1000);
  };

  ["load", "mousemove", "keypress", "click", "scroll"].forEach(evt =>
    window.addEventListener(evt, resetTimer)
  );

  resetTimer();

  return () => clearTimeout(inactivityTimer);
};
