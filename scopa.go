package main

import (
	"fmt"
	"math/rand"
	"time"//a fim de que o jogador possa ler o resultado
)

type carta struct{
    naipe string;
    valor int;//1, 2, 3, 4, 5, 6, 7, 8(fante), 9(cavalo) ,10(rei)
}

func novacarta(naipe string, valor int) carta{
	c:=carta{naipe:naipe, valor:valor};
	return c;
}

func roboescolha(mesa []carta, mao []carta, nivel int)int {
	var pontcarta []int;
	for i:=0; i<len(mao); i++{
		pontcarta=append(pontcarta, 0);
	}
	var n int=0;//qual carta
	for i:=0; i<len(pontcarta); i++{
		pontcarta[i]=0;
	}
	if nivel==0{//tenta sempre deitar
		for i:=0; i<len(mao); i++{
			for j:=0; j<len(mesa); j++{
				if mesa[j].valor==mao[i].valor{
					pontcarta[i]-=1;
				}
			}
			for j:=0; j<len(mesa)-1; j++{
				for k:=j+1; k<len(mesa); k++{
					if mesa[j].valor+mesa[k].valor==mao[i].valor{
						pontcarta[i]-=2;
					}
				}
			}
			for j:=0; j<len(mesa)-2; j++{
				for k:=j+1; k<len(mesa)-1; k++{
					for h:=k+1; h<len(mesa); h++{
						if mesa[j].valor+mesa[k].valor+mesa[h].valor==mao[i].valor{
							pontcarta[i]-=3;
						}
					}
				}
			}
		}
		//teoricamente a que tiver o menor número será a melhor
		for i:=0; i<len(pontcarta); i++{
			if pontcarta[i]>pontcarta[n]{
				n=i;
			}
		}
		return n;
	}else if nivel==1{//aleatório
		return rand.Intn(len(mao));
	}else if nivel==2{//pega a primeira
		return 0;
	}else{//tenta sempre capturar
		for i:=0; i<len(mao); i++{
			for j:=0; j<len(mesa); j++{
				if mesa[j].valor==mao[i].valor{
					pontcarta[i]+=1;
				}
			}
			for j:=0; j<len(mesa)-1; j++{
				for k:=j+1; k<len(mesa); k++{
					if mesa[j].valor+mesa[k].valor==mao[i].valor{
						pontcarta[i]+=2;
					}
				}
			}
			for j:=0; j<len(mesa)-2; j++{
				for k:=j+1; k<len(mesa)-1; k++{
					for h:=k+1; h<len(mesa); h++{
						if mesa[j].valor+mesa[k].valor+mesa[h].valor==mao[i].valor{
							pontcarta[i]+=3;
						}
					}
				}
			}
		}
		//teoricamente a que tiver o menor número será a melhor
		for i:=0; i<len(pontcarta); i++{
			if pontcarta[i]>pontcarta[n]{
				n=i;
			}
		}
		return n;
	}
}

func pontosmn(a int, b int) int{
	if b==7{
		return b;
	}
	//a partir daqui b!=7
	if a==7{
		return a;
	}else if a==6{
		return a;
	}else if a==1&&(((b<6)&&(b>1))||b>7){
		return a;
	}else if a==5&&(((b<6)&&(b>1))||b>7){
		return a;
	}else if a==4&&(((b<5)&&(b>1))||b>7){
		return a;
	}else if a==3&&(((b<4)&&(b>1))||b>7){
		return a;
	}else if a==2&&(b==2||b>7){
		return a;
	}else{
		return b;
	}
}
func main(){
	//deleta:="\033[2J";
	var vez int=0;//se par, vez do jogador; se impar, vez do computador 
	var baralho [] carta;//vai ter 40
	var naipes = [4] string{"denara", "paus", "espadas", "copas"};
	var a int=-1;
	//cria o baralho
	for i:=0; i<4; i++{
		for j:=1; j<11; j++{
			a++;
			baralho=append(baralho, novacarta(naipes[i],j));
		}
	}


	//embaralha o baralho
	genesis:
	var troca carta;
	var b int;
	var c int;
	for i:=0; i<80; i++{
		b=rand.Intn(40);
		c=rand.Intn(40);
		troca=baralho[b];
		baralho[b]=baralho[c];
		baralho[c]=troca;
	}
	//cria as mãos e mesa e as preenche
	var mesa[]carta;
	var maojogador[]carta;//cada mão tem no máximo, três cartas
	var maocomputador[]carta;
	var escopasjogador int;
	var escopascomputador int;
	var montejogador[]carta;//conte o monte
	var montecomputador[]carta;
	var escolhajogador int;
	var escolhacomputador int;

	mesa=baralho[:4];
	baralho=baralho[4:];
	qtdreis:=0;
	for i:=0; i<len(mesa); i++{
		if mesa[i].valor==10{
			qtdreis++;
		}
	}
	if qtdreis>=3{
		goto genesis;
	}
	quercomecar:=true;
	nivelcomputador:=0;

	fmt.Println("Quer começar? [0:não, 1:sim]");
	fmt.Scanln(&quercomecar);
	fmt.Println("Contra que robô você quer jogar?[0 a 2 - fácil, 3 - médio]");
	fmt.Scanln(&nivelcomputador);
	if quercomecar{
		vez=0;
	}else{
		vez=1;
	}
	for len(baralho)>=0{//enquanto houver cartas distribuíveis(não nulas)
		distribuicartas:
			if len(baralho)==0{
				if (vez-1)%2==0{
					for i:=0; i<len(mesa); i++{
						montejogador=append(montejogador, mesa[i]);
					}
					goto contagem;
				}else{
					for i:=0; i<len(mesa); i++{
						montecomputador=append(montecomputador, mesa[i]);
					}
					goto contagem;
				}
			}
			maojogador=baralho[:3];//pensei em usar um append, mas nâo é necessário; se for fazer devidas alterações
			baralho=baralho[3:];
			maocomputador=baralho[:3];
			baralho=baralho[3:];
		for (len(maojogador)>0&&len(maocomputador)>=0)||(len(maojogador)>=0&&len(maocomputador)>0){
			l:=0;
			for i:=0; i<len(mesa); i++{
				if mesa[i].naipe=="denara"{
					troca=mesa[i];
					mesa[i]=mesa[l];
					mesa[l]=troca;
					l++;
				}
			}
			/*
			//imprime as cartas
			//fmt.Print(string(deleta));
			fmt.Println("");
			fmt.Println("")
			fmt.Print("		");
			for i:=0; i<len(maocomputador); i++{
				fmt.Print("[] ");
			}
			//fmt.Print(maocomputador);//provisorio
			fmt.Println("\n");
			for i:=0; i<len(mesa); i++{
				fmt.Print("[");
				fmt.Print(mesa[i].valor);
				fmt.Print(" de ");
				fmt.Print(mesa[i].naipe);
				fmt.Print("]   ");
			}
			//fmt.Println(mesa);
			fmt.Println("");
			fmt.Println("");
			//fmt.Println(maojogador);
			fmt.Print("   ");
			for i:=0; i<len(maojogador); i++{
				fmt.Print(i);
				fmt.Print(":");
				fmt.Print("[");
				fmt.Print(maojogador[i].valor);
				fmt.Print(" de ");
				fmt.Print(maojogador[i].naipe);
				fmt.Print("]");
				fmt.Print("  ");
			}
			fmt.Println("\n");
			*/
			//goto fim;
			//jogo
			if vez%2==0{
			//imprime as cartas
			fmt.Println("");
			fmt.Println("")
			fmt.Print("		");
			for i:=0; i<len(maocomputador); i++{
				fmt.Print("[] ");
			}
			fmt.Println("\n");
			for i:=0; i<len(mesa); i++{
				fmt.Print("[");
				fmt.Print(mesa[i].valor);
				fmt.Print(" de ");
				fmt.Print(mesa[i].naipe);
				fmt.Print("]   ");
			}
			//fmt.Println(mesa);
			fmt.Println("");
			fmt.Println("");
			//fmt.Println(maojogador);
			fmt.Print("   ");
			for i:=0; i<len(maojogador); i++{
				fmt.Print(i);
				fmt.Print(":");
				fmt.Print("[");
				fmt.Print(maojogador[i].valor);
				fmt.Print(" de ");
				fmt.Print(maojogador[i].naipe);
				fmt.Print("]");
				fmt.Print("  ");
			}
			fmt.Println("\n");
				coletadedados:
				fmt.Scanln(&escolhajogador);
				if escolhajogador>=len(maojogador)||escolhajogador<0{
					goto coletadedados;
				}
				//fmt.Print("tu:");
				//fmt.Println(escolhajogador);
				for i:=0; i<len(mesa); i++{//se houver uma carta de valor identico na mesa
					if mesa[i].valor==maojogador[escolhajogador].valor{
						if len(mesa)-1==0{
							escopasjogador++;
							fmt.Println("Escopa do jogador!");
						}
						troca=maojogador[escolhajogador];
						maojogador[escolhajogador]=maojogador[0];
						maojogador[0]=troca;//coloca carta no começo pela praticidade
						//coloca cartas nos devidos lugares
						montejogador=append(montejogador, maojogador[0]);
						montejogador=append(montejogador, mesa[i]);
						maojogador=maojogador[1:];
						troca=mesa[i];
						mesa[i]=mesa[0];
						mesa[0]=troca;
						mesa=mesa[1:];
						
						goto vez;
					}
				}
				for i:=0; i<len(mesa)-1; i++{//se houver mais de uma carta que somada dá a escolha
					for j:=i+1; j<len(mesa); j++{
						if mesa[i].valor+mesa[j].valor==maojogador[escolhajogador].valor{
							if len(mesa)-2==0{
								escopasjogador++;
								fmt.Println("Escopa do jogador!");
							}
							//fmt.Println(i);fmt.Println(j);
							troca=mesa[i];
							mesa[i]=mesa[0];
							mesa[0]=troca;
							montejogador=append(montejogador, mesa[0]);
							//mesa=mesa[1:];//carta 1 da mesa
							troca=mesa[j];
							mesa[j]=mesa[1];
							mesa[1]=troca;
							montejogador=append(montejogador, mesa[1]);
							mesa=mesa[2:];//carta 2 da mesa
							//fmt.Println("!!!!!!!!!!!!!!!!");
							troca=maojogador[escolhajogador];
							maojogador[escolhajogador]=maojogador[0];
							maojogador[0]=troca;
							montejogador=append(montejogador, maojogador[0]);
							maojogador=maojogador[1:];
							goto vez;
						}
					}
				}
				if len(mesa)>=3{
					for i:=0; i<len(mesa)-2; i++{//se houver mais de uma carta que somada dá a escolha
						for j:=i+1; j<len(mesa)-1; j++{
							for k:=j+1; k<len(mesa); k++{
								if mesa[i].valor+mesa[j].valor+mesa[k].valor==maojogador[escolhajogador].valor{
									if len(mesa)-3==0{
										escopasjogador++;
										fmt.Println("Escopa do jogador!");
									}
									//fmt.Println(i);fmt.Println(j);
									troca=mesa[i];
									mesa[i]=mesa[0];
									mesa[0]=troca;
									montejogador=append(montejogador, mesa[0]);
									//mesa=mesa[1:];//carta 1 da mesa
									troca=mesa[j];
									mesa[j]=mesa[1];
									mesa[1]=troca;
									montejogador=append(montejogador, mesa[1]);
									//
									troca=mesa[k];
									mesa[k]=mesa[2];
									mesa[2]=troca;
									montejogador=append(montejogador, mesa[2]);
									mesa=mesa[3:];//carta 3 da mesa
									//fmt.Println("!!!!!!!!!!!!!!!!");
									troca=maojogador[escolhajogador];
									maojogador[escolhajogador]=maojogador[0];
									maojogador[0]=troca;
									montejogador=append(montejogador, maojogador[0]);
									maojogador=maojogador[1:];
									goto vez;
								}
							}
						}
					}
				}
				if len(mesa)>=4{//só coloquei, porque joguei e o computador não contou minha escopa de quatro
					for i:=0; i<len(mesa)-3; i++{//se houver mais de uma carta que somada dá a escolha
						for j:=i+1; j<len(mesa)-2; j++{
							for k:=j+1; k<len(mesa)-1; k++{
								for g:=k+1; g<len(mesa); g++{
									if mesa[i].valor+mesa[j].valor+mesa[k].valor+mesa[g].valor==maojogador[escolhajogador].valor{
										if len(mesa)-4==0{
											escopasjogador++;
											fmt.Println("Escopa do jogador!");
										}
										//fmt.Println(i);fmt.Println(j);
										troca=mesa[i];
										mesa[i]=mesa[0];
										mesa[0]=troca;
										montejogador=append(montejogador, mesa[0]);
										//mesa=mesa[1:];//carta 1 da mesa
										troca=mesa[j];
										mesa[j]=mesa[1];
										mesa[1]=troca;
										montejogador=append(montejogador, mesa[1]);
										//
										troca=mesa[k];
										mesa[k]=mesa[2];
										mesa[2]=troca;
										montejogador=append(montejogador, mesa[2]);
										//
										troca=mesa[g];
										mesa[g]=mesa[3];
										mesa[3]=troca;
										montejogador=append(montejogador, mesa[3]);
				
										mesa=mesa[4:];//carta 4 da mesa
										//fmt.Println("!!!!!!!!!!!!!!!!");
										troca=maojogador[escolhajogador];
										maojogador[escolhajogador]=maojogador[0];
										maojogador[0]=troca;
										montejogador=append(montejogador, maojogador[0]);
										maojogador=maojogador[1:];
										goto vez;
									}
								}
							}
						}
					}
				}
				//se não houver somas de cartas na mesa que se igualem à carta, a carta vai à mesa
				troca=maojogador[escolhajogador];
				maojogador[escolhajogador]=maojogador[0];
				maojogador[0]=troca;//coloca carta no começo pela praticidade
				mesa=append(mesa, maojogador[0]);
				maojogador=maojogador[1:];
				
				vez:
				vez++;
				//jogador
			}else{//fazer algoritmo inteligente
				//algoritmo computador burro
				//escolhacomputador=rand.Intn(len(maocomputador));
				escolhacomputador=roboescolha(mesa, maocomputador, nivelcomputador);
				for i:=0; i<len(mesa); i++{//se houver uma carta de valor identico na mesa
					if mesa[i].valor==maocomputador[escolhacomputador].valor{
						if len(mesa)-1==0{
							escopascomputador++;
							fmt.Println("Escopa do computador!");
						}
						troca=maocomputador[escolhacomputador];
						maocomputador[escolhacomputador]=maocomputador[0];
						maocomputador[0]=troca;//coloca carta no começo pela praticidade
						//coloca cartas nos devidos lugares
						montecomputador=append(montecomputador, maocomputador[0]);
						montecomputador=append(montecomputador, mesa[i]);
						maocomputador=maocomputador[1:];
						troca=mesa[i];
						mesa[i]=mesa[0];
						mesa[0]=troca;
						mesa=mesa[1:];
						
						goto vezc;
					}
				}
				for i:=0; i<len(mesa)-1; i++{//se houver mais de uma carta que somada dá a escolha
					for j:=i+1; j<len(mesa); j++{
						if mesa[i].valor+mesa[j].valor==maocomputador[escolhacomputador].valor{
							if len(mesa)-2==0{
								escopascomputador++;
								fmt.Println("Escopa do computador!");
							}
							//fmt.Println(i);fmt.Println(j);
							troca=mesa[i];
							mesa[i]=mesa[0];
							mesa[0]=troca;
							montecomputador=append(montecomputador, mesa[0]);
							//mesa=mesa[1:];//carta 1 da mesa
							troca=mesa[j];
							mesa[j]=mesa[1];
							mesa[1]=troca;
							montecomputador=append(montecomputador, mesa[1]);
							mesa=mesa[2:];//carta 2 da mesa
							//fmt.Println("!!!!!!!!!!!!!!!!");
							troca=maocomputador[escolhacomputador];
							maocomputador[escolhacomputador]=maocomputador[0];
							maocomputador[0]=troca;
							montecomputador=append(montecomputador, maocomputador[0]);
							maocomputador=maocomputador[1:];
							goto vezc;
						}
					}
				}
				if len(mesa)>=3{
					for i:=0; i<len(mesa)-2; i++{//se houver mais de uma carta que somada dá a escolha
						for j:=i+1; j<len(mesa)-1; j++{
							for k:=j+1; k<len(mesa); k++{
								if mesa[i].valor+mesa[j].valor+mesa[k].valor==maocomputador[escolhacomputador].valor{
									if len(mesa)-3==0{
										escopascomputador++;
										fmt.Println("Escopa do computador!");
									}
									//fmt.Println(i);fmt.Println(j);
									troca=mesa[i];
									mesa[i]=mesa[0];
									mesa[0]=troca;
									montecomputador=append(montecomputador, mesa[0]);
									//mesa=mesa[1:];//carta 1 da mesa
									troca=mesa[j];
									mesa[j]=mesa[1];
									mesa[1]=troca;
									montecomputador=append(montecomputador, mesa[1]);
									//
									troca=mesa[k];
									mesa[k]=mesa[2];
									mesa[2]=troca;
									montecomputador=append(montecomputador, mesa[2]);
									mesa=mesa[3:];//carta 3 da mesa
									//fmt.Println("!!!!!!!!!!!!!!!!");
									troca=maocomputador[escolhacomputador];
									maocomputador[escolhacomputador]=maocomputador[0];
									maocomputador[0]=troca;
									montecomputador=append(montecomputador, maocomputador[0]);
									maocomputador=maocomputador[1:];
									goto vezc;
								}
							}
						}
					}
				}
				if len(mesa)>=4{
					for i:=0; i<len(mesa)-3; i++{//se houver mais de uma carta que somada dá a escolha
						for j:=i+1; j<len(mesa)-2; j++{
							for k:=j+1; k<len(mesa)-1; k++{
								for g:=k+1; g<len(mesa); g++{
									if mesa[i].valor+mesa[j].valor+mesa[k].valor+mesa[g].valor==maocomputador[escolhacomputador].valor{
										if len(mesa)-4==0{
											escopascomputador++;
											fmt.Println("Escopa do computador!");
										}
										//fmt.Println(i);fmt.Println(j);
										troca=mesa[i];
										mesa[i]=mesa[0];
										mesa[0]=troca;
										montecomputador=append(montecomputador, mesa[0]);
										//mesa=mesa[1:];//carta 1 da mesa
										troca=mesa[j];
										mesa[j]=mesa[1];
										mesa[1]=troca;
										montecomputador=append(montecomputador, mesa[1]);
										//
										troca=mesa[k];
										mesa[k]=mesa[2];
										mesa[2]=troca;
										montecomputador=append(montecomputador, mesa[2]);
										//
										troca=mesa[g];
										mesa[g]=mesa[3];
										mesa[3]=troca;
										montecomputador=append(montecomputador, mesa[3]);
				
										mesa=mesa[4:];//carta 4 da mesa
										//fmt.Println("!!!!!!!!!!!!!!!!");
										troca=maocomputador[escolhacomputador];
										maocomputador[escolhacomputador]=maocomputador[0];
										maocomputador[0]=troca;
										montecomputador=append(montecomputador, maocomputador[0]);
										maocomputador=maocomputador[1:];
										goto vezc;
									}
								}
							}
						}
					}
				}
				//se não houver somas de cartas na mesa que se igualem à carta, a carta vai à mesa
				troca=maocomputador[escolhacomputador];
				maocomputador[escolhacomputador]=maocomputador[0];
				maocomputador[0]=troca;//coloca carta no começo pela praticidade
				mesa=append(mesa, maocomputador[0]);
				maocomputador=maocomputador[1:];
				
				vezc:
				vez++;
			}
		}
		goto distribuicartas;//quando acabarem as cartas, distribui-las é o correto.
	}
	contagem:
	pontosj:=0;
	pontosc:=0;
	setebelo:=false;//true-jogador; false-computador
	ouroj:=0;
	ouroc:=0;
	//Quem tem o maior monte?
	fmt.Print("Jogador: ");
	fmt.Print(len(montejogador));
	fmt.Println(" cartas.");
	fmt.Print("Computador: ");
	fmt.Print(len(montecomputador));
	fmt.Println(" cartas.");
	fmt.Print("Jogador: ");
	fmt.Print(escopasjogador);
	fmt.Println(" escopas.");
	fmt.Print("Computador: ");
	fmt.Print(escopascomputador);
	fmt.Println(" escopas.");

	if len(montejogador)>len(montecomputador){
		pontosj++;
	}else if len(montejogador)<len(montecomputador){
		pontosc++;
	}else{
		pontosj+=0;
	}
	//Quem capturou o sete de ouros
	for i:=0; i<len(montejogador); i++{
		if montejogador[i].naipe=="denara"&&montejogador[i].valor==7{
			fmt.Println("Jogador capturou o sete de ouros.");
			setebelo=true;
		}
	}
	if setebelo{
		pontosj++;
	}else{
		fmt.Println("Computador capturou o sete de ouros.");
		pontosc++;
	}
	//fmt.Println("Quem capturou mais cartas do naipe de ouros?");
	for i:=0; i<len(montejogador); i++{
		if montejogador[i].naipe=="denara"{
			ouroj++;
		}
	}
	for i:=0; i<len(montecomputador); i++{
		if montecomputador[i].naipe=="denara"{
			ouroc++;
		}
	}
	if ouroj>ouroc{
		pontosj++;
		fmt.Print("Jogador: ");
		fmt.Print(ouroj);
		fmt.Println(" cartas do naipe de ouros");
	}else if ouroc>ouroj{
		pontosc++;
		fmt.Print("Computador: ");
		fmt.Print(ouroc);
		fmt.Println(" cartas do naipe de ouros");
	}else{
		pontosj+=0;
	}
	//fmt.Println("Quem teve o melhor número?");//"primiera" ou algo do tipo
	mnj:=[4]int {0, 0, 0, 0};//ouro, pau, espada, copa 
	mnc:=[4]int {0, 0, 0, 0};
	for i:=0; i<len(montejogador); i++{
		switch(montejogador[i].naipe){
			case "denara":
				if pontosmn(montejogador[i].valor, mnj[0])==montejogador[i].valor{//montejogador[i].valor>mnj[0]{
					mnj[0]=montejogador[i].valor;
				}
			case "paus":
				if pontosmn(montejogador[i].valor, mnj[1])==montejogador[i].valor{//montejogador[i].valor>mnj[1]{
					mnj[1]=montejogador[i].valor;
				}
			case "espadas":
				if pontosmn(montejogador[i].valor, mnj[2])==montejogador[i].valor{//montejogador[i].valor>mnj[2]{
					mnj[2]=montejogador[i].valor;
				}
			case "copas":
				if pontosmn(montejogador[i].valor, mnj[3])==montejogador[i].valor{//montejogador[i].valor>mnj[3]{
					mnj[3]=montejogador[i].valor;
				}
		}
	}
	for i:=0; i<len(montecomputador); i++{
		switch(montecomputador[i].naipe){
			case "denara":
				if pontosmn(montecomputador[i].valor, mnj[0])==montecomputador[i].valor{//montecomputador[i].valor>mnj[0]{
					mnj[0]=montecomputador[i].valor;
				}
			case "paus":
				if pontosmn(montecomputador[i].valor, mnj[1])==montecomputador[i].valor{//montecomputador[i].valor>mnj[1]{
					mnj[1]=montecomputador[i].valor;
				}
			case "espadas":
				if pontosmn(montecomputador[i].valor, mnj[2])==montecomputador[i].valor{//montecomputador[i].valor>mnj[2]{
					mnj[2]=montecomputador[i].valor;
				}
			case "copas":
				if pontosmn(montecomputador[i].valor, mnj[3])==montecomputador[i].valor{//montecomputador[i].valor>mnj[3]{
					mnj[3]=montecomputador[i].valor;
				}
		}
	}
	for i:=0; i<len(mnj); i++{//os dois fors são obrigatórios (consulte a wikipédia)
		switch(mnj[i]){
		case 7:
			mnj[i]=21;
		case 6:
			mnj[i]=18;
		case 1:
			mnj[i]=16;
		case 5:
			mnj[i]=15;
		case 4:
			mnj[i]=14;
		case 3:
			mnj[i]=13;
		case 2:
			mnj[i]=12;
		default:
			mnj[i]=10;
		}
	}
	for i:=0; i<len(mnc); i++{
		switch(mnc[i]){
		case 7:
			mnc[i]=21;
		case 6:
			mnc[i]=18;
		case 1:
			mnc[i]=16;
		case 5:
			mnc[i]=15;
		case 4:
			mnc[i]=14;
		case 3:
			mnc[i]=13;
		case 2:
			mnc[i]=12;
		default:
			mnc[i]=10;
		}
	}
	if mnj[0]+mnj[1]+mnj[2]+mnj[3]>mnc[0]+mnc[1]+mnc[2]+mnc[3]{
		pontosj++;
		fmt.Println("Jogador obteve o melhor número.");
	}else if mnj[0]+mnj[1]+mnj[2]+mnj[3]<mnc[0]+mnc[1]+mnc[2]+mnc[3]{
		pontosc++;
		fmt.Println("Computador obteve o melhor número.");
	}else{
		fmt.Println("Ambos obtiveram o mesmo \"melhor número\".");
	}
	pontosj+=escopasjogador;
	pontosc+=escopascomputador;
	if pontosc<pontosj{
		fmt.Println("Jogador ganhou.");
	}else if pontosc>pontosj{
		fmt.Println("Computador ganhou.");
	}
	time.Sleep(20*time.Second);
	//verificações
	/*
	for i:=0; i<len(baralho); i++{
		fmt.Println(baralho[i]);
	}
	fmt.Print("CARTAS NA MESA: ");
	fmt.Println(mesa);
	fmt.Print("QUANTAS CARTAS HÁ NO BARALHO: ");
	fmt.Println(len(baralho));
	fmt.Println("JOGADOR: ");
	fmt.Println(maojogador);
	fmt.Println("COMPUTADOR: ");
	fmt.Println(maocomputador);
	*/
}