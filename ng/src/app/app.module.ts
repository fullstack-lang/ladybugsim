import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { LadybugsimModule } from 'ladybugsim'

// angular split
import { AngularSplitModule } from 'angular-split';

// gongsim stack
import { GongsimcontrolModule } from 'gongsimcontrol'
import { GongsimModule } from 'gongsim'

// mandatory
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,

    // gongsim stack
    GongsimcontrolModule,
    GongsimModule,

    // angulat split
    AngularSplitModule,

    HttpClientModule,
    LadybugsimModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
