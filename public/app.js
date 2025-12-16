import { API } from './services/API.js';
import './components/HomePage.js';


window.app = {
    search: event => {
        event.preventDefault();
        const query = document.getElementById('search-input').value;
    },
    
    api: API
}