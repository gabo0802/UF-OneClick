import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

//Angular Material imports
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,    
  ],
  exports: [
    MatSlideToggleModule,
    MatToolbarModule,
    MatButtonModule
  ]
})
export class MaterialDesignModule { }
