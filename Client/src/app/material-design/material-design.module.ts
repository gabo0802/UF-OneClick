import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

//Angular Material imports
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import {  MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatDialogModule } from '@angular/material/dialog';
import { ScrollingModule } from '@angular/cdk/scrolling';
import { MatCarouselModule } from '@magloft/material-carousel';
@NgModule({
  declarations: [],
  imports: [
    CommonModule,    
    MatCarouselModule.forRoot()
  ],
  exports: [
    MatSlideToggleModule,
    MatToolbarModule,
    MatButtonModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatMenuModule,
    MatDialogModule,
    ScrollingModule,
  ]
})
export class MaterialDesignModule { }
