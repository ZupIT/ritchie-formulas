import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { Beagle } from './beagle.module';

import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    Beagle
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
