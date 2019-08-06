import {HttpClient} from '@angular/common/http';
import {Router} from '@angular/router';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs/internal/Observable';
declare var $:any;

@Injectable()
export class Service {
  constructor(private http: HttpClient,
              private router: Router) {
  }

  getImagesList(): Observable<string[]> {
    return this.http.get<string[]>('/images/list');
  }

  syncGet(url: string, success: (resp: any) => void) {
    $.ajax({
      url: url,
      type: 'GET',
      dataType: 'JSON',
      success: success,
      async: false
    });
  }
}
