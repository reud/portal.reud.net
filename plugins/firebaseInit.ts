import firebase from 'firebase';

if (!firebase.apps.length) {
  firebase.initializeApp({
    apiKey: 'AIzaSyBAq_lidTR4FtzF1SCLjDGl3QAisLXhbF8',
    authDomain: 'portal-apps.firebaseapp.com',
    databaseURL: 'https://portal-apps.firebaseio.com',
    projectId: 'portal-apps',
    storageBucket: 'portal-apps.appspot.com',
    messagingSenderId: '823967458159',
    appId: '1:823967458159:web:adc6459ef657fb5ae8fe9a',
    measurementId: 'G-JTSLKXBQGJ'
  });
}

export default firebase;
