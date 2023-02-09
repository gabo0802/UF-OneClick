import { Component, ViewChild, ElementRef } from '@angular/core';
//import { MouseEvent } from  '@angular/platform-browser/src/dom/events/event';
import { fromEvent, interval, Subject } from 'rxjs';
import { map, takeUntil } from 'rxjs/operators';

@Component({
  selector: 'app-scroll-bar',
  template: `
    <mat-toolbar color="primary" class="scroll-bar">
      <div class="scroll-area" #scrollArea>
        <ng-container *ngFor="let logo of logos">
          <img [src]="logo" class="logo" />
        </ng-container>
      </div>
    </mat-toolbar>
  `,
  styles: [`
    .scroll-bar {
      height: 20%;
      width: 100%;
      display: flex;
      justify-content: flex-start;
      overflow-x: scroll;
    }

    .scroll-area {
      display: flex;
    }

    .logo {
      height: 100%;
      margin-right: 16px;
    }
  `]
})
export class ScrollBarComponent {
  logos = [
    'https://logo.clearbit.com/netflix.com',
    'https://logo.clearbit.com/amazon.com',
    'https://logo.clearbit.com/spotify.com',
    'https://logo.clearbit.com/hulu.com',
    'https://logo.clearbit.com/disneyplus.com',
  ];

  @ViewChild('scrollArea') scrollArea: ElementRef = new ElementRef('scrollArea');

  ngAfterViewInit() {
    const mouseDown$ = fromEvent(this.getScrollArea(), 'mousedown');
    const mouseUp$ = fromEvent(document, 'mouseup');
    const mouseMove$ = fromEvent(document, 'mousemove');
      /*
    mouseDown$.subscribe((event: MouseEvent) => {
      const startX = (event as MouseEvent).clientX;

      const scroll$ = mouseMove$.pipe(
        map(moveEvent => startX - (moveEvent as MouseEvent).clientX),
        takeUntil(mouseUp$),
      );

      scroll$.subscribe(offset => {
        this.getScrollArea().scrollLeft += offset;
      });
    }); */
  }

  private getScrollArea(): HTMLElement {
    return this.scrollArea.nativeElement;
  } 
}