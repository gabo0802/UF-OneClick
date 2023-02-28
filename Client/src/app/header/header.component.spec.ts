import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { MaterialDesignModule } from '../material-design/material-design.module';
import { AuthService } from '../auth.service';
import { HeaderComponent } from './header.component';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        HeaderComponent
       ],
       imports: [
        MaterialDesignModule, 
        RouterTestingModule
       ],
       providers: [
        AuthService
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('if user is not logged in, redirected on header image click to "/" ', () => {    

    expect(component.changeRoute()).toEqual(['/'])
  })
});
