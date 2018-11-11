import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-page',
  templateUrl: './page.component.html',
  styleUrls: ['./page.component.css']
})
export class PageComponent implements OnInit {

  page = {
    title: 'Home',
    subtitle: 'Welcome Home',
    content: 'Some content'
  };
  constructor() { }

  ngOnInit() {
  }

  redirectToSpotify() {
    window.location.href = "https://accounts.spotify.com/authorize?client_id=725c53094c0449379beb24431dda70cf&response_type=code&redirect_uri=http%3A%2F%2F467cf652.ngrok.io%2Fgenres"
  }
}
