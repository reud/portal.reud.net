// @ts-ignore
import firebase from '~/plugins/firebaseInit';

export const fetchUser = (): Promise<firebase.User | boolean> => {
  return new Promise((resolve) => {
    firebase.auth().onAuthStateChanged((user: firebase.User | null) => {
      resolve(user || false);
    });
  });
};

export const handleGoogleLogin = () => {
  const googleAuthProvider = new firebase.auth.GoogleAuthProvider();
  return firebase.auth().signInWithPopup(googleAuthProvider);
};
// @ts-ignore
