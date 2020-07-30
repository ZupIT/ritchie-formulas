import { Component } from '@angular/core';
import { LoadParams } from '@zup-it/beagle-web';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  title = 'beagleWeb';

  loadParams: LoadParams;

  constructor() {
    this.loadParams = {
      path: '/test.json'
    };
  }
}
