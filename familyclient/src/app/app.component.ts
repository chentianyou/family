import {Component, OnInit} from '@angular/core';
import {Service} from './common/service/service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  images: any[];

  constructor(
    private service: Service
  ) {
  }

  ngOnInit() {
    this.images = [];
    this.service.syncGet('/images/list', reps => {
      reps.forEach(img => {
        this.images.push({source: '/images/' + img, alt: 'my son', title: img});
      });
    });
  }
}
