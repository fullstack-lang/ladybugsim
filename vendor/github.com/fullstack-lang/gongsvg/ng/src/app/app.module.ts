import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { GongsvgModule } from 'gongsvg'
import { GongsvgspecificModule } from 'gongsvgspecific'

// mandatory
import { HttpClientModule } from '@angular/common/http';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    AngularSplitModule,

    HttpClientModule,
    GongsvgModule,
    GongsvgspecificModule

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
