import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ServerResponse } from '../models/server-response.model';
import axios from 'axios';
import { redirectUri } from '../common';

const BASE_GENRES_ENDPOINT = 'http://localhost:8000/api/genres';

@Component({
  selector: 'app-genre',
  templateUrl: './genre.component.html',
  styleUrls: ['./genre.component.css']
})
export class GenreComponent implements OnInit {
  code: string;
  genres: string[];

  constructor(private route: ActivatedRoute) {}

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.code = params.code;
      this.getGenres();
    });
  }

  private getGenres() {
    axios.request<ServerResponse<string>>({
      url: this.buildGenresUrl()
    }).then(resp => {
      console.log(resp);
      this.genres = [];
    });
  }

  private buildGenresUrl(): string {
    return BASE_GENRES_ENDPOINT + '?code=' + this.code + '&redirect_uri=' + redirectUri;
  }
}
