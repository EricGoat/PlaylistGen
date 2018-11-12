import { Component, OnInit } from '@angular/core';
import axios from 'axios';
import { ServerResponse } from '../models/server-response.model';
import { redirectUri } from '../common';

const BASE_AUTHORIZE_URL = 'https://accounts.spotify.com/authorize';

@Component({
  selector: 'app-page',
  templateUrl: './page.component.html',
  styleUrls: ['./page.component.css']
})
export class PageComponent implements OnInit {
  clientId: string;

  constructor() { }

  ngOnInit() {
    axios.request<ServerResponse<string>>({
      url: 'http://localhost:8000/api/client_id'
    }).then(resp => {
      this.clientId = resp.data.data;
    });
  }

  redirectToSpotify() {
    window.location.href = this.buildAuthorizationUrl();
  }

  private buildAuthorizationUrl(): string {
    return BASE_AUTHORIZE_URL + '?client_id=' + this.clientId + '&response_type=code&redirect_uri=' + redirectUri;
  }
}
